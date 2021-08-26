# Technical Assessment Go
This software is a Go HTTP server that proxies requests for PNG and JPEG XYZ map tiles to a configurable origin domain (i.e. https://maps.wikimedia.org/) and grayscales the response images on-the-fly.  It is a prototype that is part of a technical assessment to demonstrate competency in Go programming, HTTP API design, and basic map tile image processing.

## Requirements
- git >= 2.30.1
- Go >= 1.16.3

## Installation, Tests, and Running the Development Server
1. Clone this GitHub repository:
```
git clone git@github.com:asonnenschein/technical-assessment-golang.git
```

2. Install project dependencies.  This command should be run at the root level of the project:
```
go get ./cmd/*
```

3. Run feature tests to confirm that the software is installed correctly.  This command should be run at the root level of the project:
```
go test internal/handlers/* 
```

4. Run development server.  This command should be run at the root level of the project:
```
go run cmd/server/main.go https://maps.wikimedia.org/
```

5. Confirm that the development server is working by visiting these URLs in your browser:
```
http://localhost:8080/osm-intl/1/0/0.png
http://localhost:8080/osm-intl/1/0/0.jpeg
```

## Notes:
- This application is a prototype, or proof of concept.  It is not intended to run in a production environment as is.
- As such, it does not currently do any application specific logging - adding DEBUG/ERROR/INFO logging would be a logical next step.
- After that, it would make sense to add more graceful failure mechanisms to the HTTP API.  While the service is generous in accepting wildcard paths, the current implementation is optimistic and is only guaranteed (i.e. tested) to work with properly formatted PNG and JPEG XYZ map tiles.
- The "WildCard" service handler is busy, i.e. it does more than one thing.  It would probably make sense to pull the image processing components out of the handler into it's own utility method that could be better unit tested.
- Speaking of tests, only a single feature test is included in this software, to demonstrate competency in automated testing.  There are no unit tests.  At a high-level, an end-to-end feature test should, in theory, cover optimistic use cases.
- The same code that is tested (handlers.go) is also the only file that includes inline documentation, to demonstrate competency in writing docstrings.
- There is a requirement in the spec to support both paletted and non-paletted images.  My understanding is that PNG is a paletted image format and JPEG is not a paletted image format.  Therefore supporting both of these content types satisfies the paletted vs. non-paletted requirement.

## Sources, Documentation, etc
- https://docs.geoserver.org/stable/en/user/tutorials/palettedimage/palettedimage.html
- https://pkg.go.dev/image?utm_source=gopls#Image
- https://stackoverflow.com/questions/8697095/how-to-read-a-png-file-in-color-and-output-as-gray-scale-using-the-go-programmin
- https://christine.website/blog/within-go-repo-layout-2020-09-07
- https://gobyexample.com/http-servers
- https://github.com/gorilla/mux
- https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
- https://www.loginradius.com/blog/async/tune-the-go-http-client-for-high-performance/
- https://blog.questionable.services/article/testing-http-handlers-go/
