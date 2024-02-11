package m3u8

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/grafov/m3u8"
)


type Segment struct {
	*m3u8.MediaSegment
	Path string
}


func newRequest(url string, headers map[string]string) (*http.Request, error) {
	fmt.Println("URL", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) ")
	req.Header.Set("accept", "*/*")
	req.Header.Set("user-agent", "curl/7.87.0")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func getM3u8ListType(url string, headers map[string]string) (m3u8.Playlist, m3u8.ListType, error) {

	req, err := newRequest(url, headers)
	if err != nil {
		return nil, 0, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()


	var byteBody []byte
	res.Body.Read(byteBody)


	fmt.Println("body", string(byteBody))

	if res.StatusCode != 200 {
		return nil, 0, errors.New(res.Status)
	}


	p, t, err := m3u8.DecodeFrom(res.Body, false)
	if err != nil {
		return nil, 0, err
	}

	return p, t, nil
}


func ParseHlsSegments(hlsURL string, headers map[string]string) ([]*Segment, error) {
	baseURL, err := url.Parse(hlsURL)
	if err != nil {
		return nil, errors.New("Invalid m3u8 url")
	}

	p, t, err := getM3u8ListType(hlsURL, headers)

	if err != nil {
		return nil, err
	}
	if t != m3u8.MEDIA {
		return nil, errors.New("No support the m3u8 format")
	}

	mediaList := p.(*m3u8.MediaPlaylist)

	fmt.Println("Media Playlist", mediaList)

	segments := []*Segment{}
	for _, seg := range mediaList.Segments {
		if seg == nil {
			continue
		}

		if !strings.Contains(seg.URI, "http") {
			segmentURL, err := baseURL.Parse(seg.URI)
			if err != nil {
				return nil, err
			}

			seg.URI = segmentURL.String()
		}

		if seg.Key == nil && mediaList.Key != nil {
			seg.Key = mediaList.Key
		}

		if seg.Key != nil && !strings.Contains(seg.Key.URI, "http") {
			keyURL, err := baseURL.Parse(seg.Key.URI)
			if err != nil {
				return nil, err
			}

			seg.Key.URI = keyURL.String()
		}

		segment := &Segment{MediaSegment: seg}
		segments = append(segments, segment)
	}

	return segments, nil
}