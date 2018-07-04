<html>
    <head>
        <title> Home </title>
        <meta http-equiv="content-type" charset="utf-8"/>
        <link rel="stylesheet" type="text/css" href="/static/css/materialize.css"/>
        <script src="/static/js/script.js"> </script>
    </head>
    <body>
        <form action="/consulta" method="POST">
            <script>
                var msg = {{.msg}}
                if(msg != ""){
                    alert(msg)
                }
            </script>
            <label> Codigo : </label>
            <input type="number" name="codigo" value="{{.codigo}}" onKeyPress="return somenteNumeros(event)" required/>
            <br/> <br/>
            <label> Nome : </label>
            <input type="text" name="nome" value="{{.nome}}" readonly="true"/>
            <br/> <br/>
            <label> Preço : </label>
            <input type="text" name="preco" value="{{.preco}}" readonly="true"/>
            <br/> <br/>
            <input type="submit" value="Consultar" class="waves-effect waves-light btn"/>
        </form>
    </body>
</html>