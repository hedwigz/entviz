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
![image](https://user-images.githubusercontent.com/8277210/129726604-90dda3b5-8e52-4a79-9017-1a106f1110b9.png)
