<html>
    <head>
        <title> Home </title>
        <meta http-equiv="content-type" charset="utf-8"/>
        <link rel="stylesheet" type="text/css" href="/static/css/materialize.css"/>
        <script src="/static/js/script.js"> </script>
    </head>
    <body>
        <script>
            var msg = {{.msg}}
            if(msg != ""){
                alert(msg)
                window.top.location = "/"
            }
        </script>
        <form action="/insere" method="POST">
            <label> Nome : </label>
            <input type="text" name="nome" class="validate" required/>
            <br/> <br/>
            <label> Preço : </label>
            <input type="text" name="preco" class="validate" onKeyPress="return somenteMoeda(event)" required/>
            <br/> <br/>
            <input type="submit" value="Cadastrar" class="waves-effect waves-light btn"/>
        </form>
    </body>
</html>