# entviz
entviz is an ent's hook that creates visual graph (html file) of your ent's schema.  
# install
```
go get github.com/hedwigz/entviz
```
add this hook to ent (see [example](examples/ent/entc.go) code)
run
```
go generate ./ent
```
your html will be saved at `ent/schema-viz.html`
# serve via http
You can use the helper function `ent.ServeEntviz` to easily serve the static html page over http
```golang
http.ListenAndServe("localhost:3002", ent.ServeEntviz())
```
# use from command line
install the cmd
```
go get github.com/hedwigz/entviz/cmd/entviz
```
then run inside your project:
```
entviz ./etc/schema
```
# example
![image (3)](https://user-images.githubusercontent.com/8277210/129726965-d3c89f1a-d66a-46b6-82a2-20f1056d350d.png)

