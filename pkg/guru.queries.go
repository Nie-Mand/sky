package pkg

var GuruGetCourseContent = `
query Course_userCourseOverview($courseId: String!) {
	userCourseOverview(courseId: $courseId) {
	  courseOverview {
		...CourseFragment
		sections {
		  ...SectionFragment
		  components {
			...ComponentFragment
		  }
		}
	  }
	}
  }
  
  fragment ComponentFragment on Component {
	id
	title
	sequence

	resources {
		title
		url
	}
  
	content {
	  ... on VideoCourseComponentContent {
		type
		
		videosources(filter: { videoType: "video/mp4" }) {
		  signedUrl
		  quality
		}
	  }
	}
  }
  
  fragment SectionFragment on Section {
	title
	sequence
  }
  
  fragment CourseFragment on CourseOverview {
	artworkUrl
	title
	description
  }
  
`



  
  
  
  
  
  