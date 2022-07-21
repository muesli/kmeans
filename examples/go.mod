module example

go 1.18

require (
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/muesli/clusters v0.0.0-20200529215643-2700303c1762
	github.com/muesli/kmeans v0.3.0
)

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/wcharczuk/go-chart/v2 v2.1.0 // indirect
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5 // indirect
)

replace github.com/muesli/kmeans => ../
