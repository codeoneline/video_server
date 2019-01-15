go build
env GOOS=linux GOARCH=amd64 go build

go install
go get
go fmt
go test
go test -v

// api design 
http status code: 
	200 get succuss
	201 create success
	204 no content
	400 page not exist
	401 no verify
	403 no auth
	500 server error
                            url                                method                      sc
* create user :            /user                                post                 201, 400, 500
* login :              /user/:username                          post                 201, 400, 500
* get user info:       /user/:username                          get                200, 400, 401, 403, 500
* del user :           /user/:username                         delete              204, 400, 401, 403, 500

* list all videos:     /user/:username/videos                   get                  200, 400, 500
* get one video:       /user/:username/videos/:vid-id           get                  200, 400, 500
* delete one video:    /user/:username/videos/:vid-id          delete              204, 400, 401, 403, 500

* show comments        /videos/:vid-id/comments                 get                  200, 400, 500
* post a comment       /videos/:vid-id/comments                 post                 201, 400, 500
* delete a comment     /videos/:vid-id/comment/:cmt-id         delete              204, 400, 401, 403, 500




handler->validation{1.request, 2.user}->business logic->response
1. data model
2. error handling