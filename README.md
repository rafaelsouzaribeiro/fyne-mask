<h1>Applying mask in golang and fyne.</h1><br />
<p>Adding mask with only numeric fields:</p><br />

```go
	myApp := app.New()
	myWindow := myApp.NewWindow("Enter the mask")

	content := di.NewDi(mask.NewMask(components.NewEnrty("Enter the mask"), "###.###.###-##", true), "Enter the mask")
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()

```
 <br />

<p>Adding mask with alphanumeric fields:</p><br />

```go
	myApp := app.New()
	myWindow := myApp.NewWindow("Enter the mask")

	content := di.NewDi(mask.NewMask(components.NewEnrty("Enter the mask"), "(##)####-######", false), "Enter the mask")
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.ShowAndRun()
 ```
