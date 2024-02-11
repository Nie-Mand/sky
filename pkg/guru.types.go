package pkg

type GuruGetCourseContentResponse struct {
	Data GuruDataType `json:"data"`
	Message string `json:"message"`
}

type GuruDataType struct {
	UserCourseOverview GuruUserCourseOverviewType `json:"userCourseOverview"`
}

type GuruUserCourseOverviewType struct {
	CourseOverview GuruCourseOverviewType `json:"courseOverview"`
}

type GuruCourseOverviewType struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Sections []GuruSectionType `json:"sections"`
}

type GuruSectionType struct {
	Title string `json:"title"`
	Sequence int `json:"sequence"`
	Components []GuruComponentType `json:"components"`
}

type GuruComponentType struct {
	Title string `json:"title"`
	Sequence int `json:"sequence"`
	Content GuruContentType `json:"content"`
	Resources []GuruResourcesType `json:"resources"`
}

type GuruContentType struct {
	Sequence int `json:"sequence"`
	Videosources []GuruVideoSourcesType `json:"videosources"`
	Type string `json:"type"`
}

type GuruVideoSourcesType struct {
	Quality string `json:"quality"`
	SignedUrl string `json:"signedUrl"`
}


type GuruResourcesType struct {
	Title string `json:"title"`
	Url string `json:"url"`
}