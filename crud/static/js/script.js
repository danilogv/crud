var cont_virgula = 0;

function somenteNumeros(e){
    var tecla = new Number();
    if(window.event) 
		tecla = e.keyCode;
    else if(e.which)
		tecla = e.which;
    else 
		return true;
    if(tecla >= 48 && tecla <= 57)
		return true;
	return false;
}

function somenteMoeda(e){
    var tecla = new Number();
    if(window.event) 
		tecla = e.keyCode;
    else if(e.which)
		tecla = e.which;
    else 
		return true;
    if(tecla >= 48 && tecla <= 57)
		return true;
	else if(tecla == 44 && cont_virgula == 0){
		cont_virgula++;
		return true;
	}
	return false;
}