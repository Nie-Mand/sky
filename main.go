package main

import (
	"Nie-Mand/sky-scraper/cmd"
	"fmt"

	"github.com/canhlinh/hlsdl"
	// "github.com/canhlinh/hlsdl"
)


func test() {

	url := "https://www.udemy.com/assets/44057352/files/2022-10-06_17-06-57-ec46d9f394ac971d33e641e57eac251d/2/hls/AVC_1024x576_1000k_AAC-HE_64k/aa0034d524974403b0cdd70c1fd8bb807479.m3u8?provider=cloudfront&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJwYXRoIjoiMjAyMi0xMC0wNl8xNy0wNi01Ny1lYzQ2ZDlmMzk0YWM5NzFkMzNlNjQxZTU3ZWFjMjUxZC8yLyIsImV4cCI6MTY4ODM1MDQxMn0.5lPLFt2JWeYF8WDdEvae3SnK7aigQsAtjKjC_5lE0YA&v=1"
	// seg, err := m3u8.ParseHlsSegments(url, nil)

	headers := map[string]string{
		"accept": "*/*",
		"user-agent": "curl/7.87.0",
	}
	
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(seg)





	// url := "https://hls-c.udemycdn.com/2022-10-06_17-06-57-ec46d9f394ac971d33e641e57eac251d/2/hls/AVC_1024x576_1000k_AAC-HE_64k/aa0034d524974403b0cdd70c1fd8bb8074790.ts?Expires=1688350413&Signature=I~MDs18n~QGHcDI3qWKBXl7xRcB51zo2gBJN1hoHPFvQYgv54wbUuK1WUJUmTx042lFk2wS0hCPbc6BiKVb-9XsN9z7J~JYB0XeKdr6aw~~a9dElbEMxuTQXvGw4mPdz4TiL5kMjMYqHT~gD0Vu1CuS15RAuFXfGDC2yPftx27osvr7PNZ-oXZQz7cXW6gVaPggoSXFFU9mFBgBSKTefELjgC6Orv8Fc7B5oljXqJTdQvjTXZqknHmkwN5878E~3uusErEwvA4EEIo-kWbA4zgLxBtbOFHVfS64~vLp9tP4BmX~jggB7sw15wyRsimI~2AUiIPw9iN3Okgmdg0lNDQ__&Key-Pair-Id=APKAITJV77WS5ZT7262A"

	// cookies := "__udmy_2_v57r=6576c02f251e4309a14ad2865e6869ef; __udmy_1_a12z_c24t=VGhlIGFuc3dlciB0byBsaWZlLCB0aGUgdW5pdmVyc2UsIGFuZCBldmVyeXRoaW5nIGlzIDQy; _pxvid=337d483e-0e83-11ee-b295-9c5572e982c4; ud_firstvisit=2023-06-19T09:25:16.667819+00:00:1qBB8O:e8G94036n_qNigd7Xa6SutstguA; ud_cache_marketplace_country=TN; ud_cache_price_country=TN; ud_cache_release=f3aae0c728f55662dffd; ud_cache_version=1; ud_cache_device=None; seen=1; __cfruid=0ccf0e8896ae776fa71d33a56e53c38f0161cf9d-1688333905; __cf_bm=5O_rEi9ZpPTCqRkakAcnBzE4nFf4O0VIaCUFtcURg1o-1688333906-0-AXNbreuUkDyFj6PfkwLZQFhF3dGIDuki+0bWTIK4Tz9ayELmfhkBvtrC4i71iS7j+vdBL0KRersvB50/fP+n1LvLLHaSDfnkMa33Fy/FNsjP; pxcts=c8049440-1920-11ee-8d7e-67564d55776d; _gid=GA1.2.1064122321.1688333911; _gcl_au=1.1.878061292.1688333914; blisspoint_fpc=c1282131-59e4-41e3-b2d9-891f2f75af99; __ssid=393fefc31adbb9015d40ff3003ae30e; _rdt_uuid=1688333914510.79aa70d5-fbf8-4f90-a710-de4b67080d4f; ki_r=; ud_user_jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MTcxMTQ3NzE4LCJlbWFpbCI6Im1vdWhpYjIwMDBAeWFob28uY29tIiwiaXNfc3VwZXJ1c2VyIjpmYWxzZSwiZ3JvdXBfaWRzIjpbXX0.A2ORI-6XDCsNhA-jcOD0MM7ASEega-51MiJMNLFgBhg; ud_credit_unseen=0; ud_credit_last_seen=None; __udmy_4_a12z=4b1e3ceb6bed3c7487ef21a7e235a5ddd99591def4633fdf5d358fe6dd40f64d; csrftoken=uxEr5PEaChNvb6xIVFhYSy0BhlETY45kaSNUk0Qsv7042PbZ59z5BwJTqBW0rNQ8; ud_locale=fr_FR; client_id=bd2565cb7b0c313f5e9bae44961e8db2; access_token=SmDcOCijkYaUWK5RfYLBAOvo4sx27fESYUhqYXSw; ud_last_auth_information=\"{\"backend\": \"udemy-auth\"\054 \"suggested_user_email\": \"mouhib2000@yahoo.com\"\054 \"suggested_user_name\": \"Chamsi\"\054 \"suggested_user_avatar\": \"https://img-c.udemycdn.com/user/50x50/anonymous_3.png\"}:1qG4q0:GqSrr_ShTycSW1fA0X1ROGT_eFc\"; dj_session_id=hxys6yvdtnw4l7acqc0vpebmix8ljjgb; ud_cache_user=171147718; ud_cache_logged_in=1; ud_cache_brand=TNfr_FR; ud_cache_campaign_code=JUST4U02223; ud_cache_language=fr; _gac_UA-12366301-1=1.1688334155.Cj0KCQjwwISlBhD6ARIsAESAmp7BaXS2VQ71hYKNc1W3pnVUkD1cfCqBoTa6f5i6LCgJyW_x3iu0DnUaAigxEALw_wcB; _gat=1; ki_s=227428%3A0.0.0.0.0; _gcl_aw=GCL.1688334160.Cj0KCQjwwISlBhD6ARIsAESAmp7BaXS2VQ71hYKNc1W3pnVUkD1cfCqBoTa6f5i6LCgJyW_x3iu0DnUaAigxEALw_wcB; _px2=eyJ1IjoiNjUyYmFiODAtMTkyMS0xMWVlLWI2OWEtZmJkNjRkMWZhYzBhIiwidiI6IjMzN2Q0ODNlLTBlODMtMTFlZS1iMjk1LTljNTU3MmU5ODJjNCIsInQiOjE2ODgzMzQ2NzMwNDUsImgiOiJhOWE4YjRhMDZlODFhNTgxZmZkYjg3NWFkODgzZThkNzI4MmY2OTlkYjA5ODA3NzJmNWQ0ZWJkZDIzMDZmZmIwIn0=; _px3=d65571c411c64b39974c5496d24da8a2f7ff9b15b5131e1f5b5d48131bf961da:M7C/VCFFWtPSRsN98RgFe03qhCx8DfrNMaUFj/zvdMgHgxiOHyMrfXrgnZIaUSqenpdqSjEiw1aU/MYSra0gTg==:1000:pF378JbGrVtnxs4GmZ+SrxGhMCEYy5vwQQH/NWLn45pxQDBJjUzTw7i4b6UzTL5Y5+hrMkwcIAbh1zGmUh+FtnpA1ag97EGQ4VxtY5YjMKMWiV166yUUsp4zghmXM4RcgEqsHD9/yLLED5iKs+a8IXDHQM+ynOODi6ncSpp0CpRS86pMpGj5xyERNbHrqbkTGy6DWPTlohJ4OWk99dFMSg==; intercom-device-id-sehj53dd=54f71673-171f-491b-9de5-e16a14c6daf6; ki_t=1688333914818%3B1688333914818%3B1688334174861%3B1%3B3; intercom-session-sehj53dd=UUoyczREVFUza0hGcmNYeWNucHk0bFpQNU42Wk5tVVdLM1IvaUFLd2FjYzJJVEVXb3lrQmM0QitsdVJaVmIvei0tQjVpb0lzRDg2WFNiNDJ5dUxpY3VKdz09--65e56c88acb03254defdeed3c808ee09baac9817; hlsJsStartLevel=5; _ga=GA1.2.857152386.1688333911; _dc_gtm_UA-12366301-1=1; _ga_7YMFEFLR6Q=GS1.1.1688333914.1.1.1688334209.0.0.0; OptanonConsent=isGpcEnabled=0&datestamp=Sun+Jul+02+2023+22%3A43%3A30+GMT%2B0100+(Central+European+Standard+Time)&version=202305.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=aa7d8b53-70dd-4e5f-8afa-f4de1c65174b&interactionCount=1&landingPath=NotLandingPage&groups=C0003%3A1%2CC0005%3A1%2CC0004%3A1%2CC0001%3A1%2CC0002%3A1&AwaitingReconsent=false; _dd_s=rum=0&expire=1688335112646; eventing_session_id=EfXZEoA_S5CyeGvI1A0vCg-1688336013167; evi=\"3@zQsIhpk18OPFxet2BDka4WQiClqWst_cC5T6GUOazi1zwIOX9jrShEju\"; ud_rule_vars=\"eJx1j91qAyEQhV9l8bbdZRx_NvosCzJ1x1aaIlE3NyHvXqGBFkJvh_Od78xNdKrv3HkP19xyL9Vbs9oImNBI1gocSU07nqxhe7KOk4-lfGYWfhK3TaRcW_9hw06dt3HfBAKqGews3QTOo_FSL7BKDfYFwANs4nWkzjTQXo74EXqllHIMrRw1crhSzfR2frTR3v4AcUQaP5Q9f_2rNIvTq0PzpKx8Obg9711nwAml18ortcjx_O_eu7h_A_20V_Y=:1qG4qz:XEiT179K_His66IqJVBgKufkLNc\""

	hlsDL := hlsdl.New(url, headers, "udemy", 64, true, "cka.mp4")
	
	filepath, err := hlsDL.Download()
	if err != nil {
		panic(err)
	}

	fmt.Println(filepath)
}



func main() {
	// test()
	cmd.Execute()
}