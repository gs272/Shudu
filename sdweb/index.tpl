<!DOCTYPE html>
<html>
  <head>
    <title>数独</title>
  </head>
  <body>
    <h1>数独</h1>
    <form action="/answer" method="post">
    <div style="width:240px">
    	{{range $i,$v:=.}}<input style="width:20px;height:20px;text-align:center" type="text" name="Box{{$i}}" value="{{$v}}">{{end}}
    </div>
     <br />
     <p>填入数独题，点击解答即可！</p>
     <input style="width:50px;height:30;text-align:center" type="submit" value="解答">
 </form>
  </body>
</html>
