$(document).ready(function() {
    

    console.log(document.cookie);
    
    $("#txtEmail").keyup(function(event) {
        if (event.keyCode === 13) {
             $("btnEnviar").click();
        }
    });
//Registro
    $("#btnEnviar").click(function() {
         var name = $("#txtTexto1").val()
         var apellidos = $("#txtTexto2").val()
         var username = $("#txtTexto3").val()
         var password = $("#txtPassword").val()
         var email = $("#txtEmail").val()
 
        console.log(name,apellidos,username, password,email );
 
        var envio = {
             name: name,
             apellidos: apellidos,
             username:username,
             password: password,
             email: email
        };
 
         $.post({
             url:"/envio",
             data: JSON.stringify(envio),
             success: function(data, status, jqXHR) {
                 console.log(data);
                 $("#txtTexto1").val('')
                 $("#txtTexto2").val('')
                 $("#txtTexto3").val('')
                 $("#txtPassword").val('')
                 $("#txtEmail").val('')
             },
             dataType: "json"
 
         }).done(function(data) {
             console.log("Petición realizada");
             if(data!=true){
                window.location.href="/loginFile";
             }
             
             //ActualizarHistorial();
         
         }).fail(function(data) {
             console.log("Petición fallida");
         
         }).always(function(data){
             console.log("Petición completa");
         });
    });
//Login
    $("#btnLogin").click(function() {
        var username = $("#txtTexto33").val()
        var password = $("#txtPassword1").val()
       console.log(username, password );

       var login = {
            username:username,
            password: password,
       };

        $.post({
            url:"/login",
            data: JSON.stringify(login),
            method:"POST",
            success: function(data, status, jqXHR) {
                console.log(data);
                
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
            if(data==true){
               window.location.href="/";
            }
           
            //ActualizarHistorial();
        
        }).fail(function(data) {
            console.log("Petición fallida");
            
        
        }).always(function(data){
            console.log("Petición completa");
        });
   });
   if (document.cookie != ""){
    $("#logeado").show();
    $("#no_logeado").hide();
    console.log('no logeado');
    
}   else{  
    $("#no_logeado").show();
    $("#logeado").hide();
    console.log('logeado');
    
}
   

   
 

})