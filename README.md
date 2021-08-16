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
your html will be save at `ent/schema-viz.html`
# example
![image](https://user-images.githubusercontent.com/8277210/129558705-69667890-cd32-416b-b66d-484a2d8d7003.png)
