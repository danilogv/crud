<html>
    <head>
        <title> Home </title>
        <meta http-equiv="content-type" charset="utf-8"/>
        <link rel="stylesheet" type="text/css" href="/static/css/materialize.css"/>
        <script src="/static/js/script.js"> </script>
    </head>
    <body>
        <form action="/altera" method="POST">
            <script>
                var msg1 = {{.msg1}}
                var msg2 = {{.msg2}}
                if(msg1 != ""){
                    alert(msg1)
                }
                else if(msg2 != ""){
                    alert(msg2)
                    window.top.location = "/"
                }
            </script>
            <label> Código : </label>
            <input type="number" name="codigo" value="{{.codigo}}" onKeyPress="return somenteNumeros(event)" required/>
            <br/> <br/>
            <label> Nome : </label>
            <input type="text" name="nome" value="{{.nome}}"/>
            <br/> <br/>
            <label> Preço : </label>
            <input type="text" name="preco" value="{{.preco}}" onKeyPress="return somenteMoeda(event)"/>
            <br/> <br/>
            <input type="submit" value="Consultar" name="botao" class="waves-effect waves-light btn"/>
            <input type="submit" value="Alterar" name="botao" class="waves-effect waves-light btn"/>
        </form>
    </body>
</html>