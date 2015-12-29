var token = "";
$(function(){

    $("#login").click(function(){
        $.ajax({
          type: "POST",
          url: "http://localhost:5000/token-auth",
          data: JSON.stringify({
            username: $("#username").val(),
            password: $("#password").val()
          }),
          success: function(o){
            if(o.TokenChanged){
                token = o.Token
            }
            console.log(arguments)
          },
          error: function(){
            console.log(arguments)
          },

          contentType: 'application/json; charset=utf-8'
        });
    });


    $("#signup").click(function(){
        $.ajax({
          type: "POST",
          url: "http://localhost:5000/signup",
          data: JSON.stringify({
            username: $("#username").val(),
            password: $("#password").val()
          }),
          success: function(o){
            if(o.TokenChanged){
                token = o.Token
            }
            console.log(arguments)
          },
          error: function(){
            console.log(arguments)
          },

          contentType: 'application/json; charset=utf-8'
        });
    });

    $("#logout").click(function(){
        $.ajax({
          type: "GET",
          url: "http://localhost:5000/logout",
          headers: {
            "Authorization" : "BEARER" + token
          },
          success: function(){
            console.log(arguments)
          },
          error: function(){
            console.log(arguments)
          },
          contentType: 'application/json; charset=utf-8'
        });
    });

    $("#refresh").click(function(){
        $.ajax({
          type: "GET",
          url: "http://localhost:5000/refresh-token-auth",
          headers: {
            "Authorization" : "BEARER" + token
          },
          success: function(o){
            if(o.TokenChanged){
                token = o.Token
            }

            console.log(arguments)
          },
          error: function(){
            console.log(arguments)
          },
          contentType: 'application/json; charset=utf-8'
        });
    });

    var conn;
    $("#connect").click(function(){


        conn = io.connect('http://localhost:5000/',{
            query: "access_token=" + token,
            upgrade: false,
            path: '/ws',
            transports: ['websocket'],
            forceNew: true
        });


        conn.on("connect_error", function(){
            console.log("Connection error");
//            conn.disconnect();
        })

        conn.on("reconnect_error", function(){
            console.log("Reconnection error");
//            conn.disconnect();
        })
    })
})