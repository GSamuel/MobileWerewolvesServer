
var looping;
var roomInfo;

$('document').ready(function(){
	$('#joinform').submit(function(e) {
	    e.preventDefault();
	    // Get all the forms elements and their values in one step
	    var code = $("#inputCode")[0].value;
	    var nickname = $("#inputNickname")[0].value
	    joinRoom(code, nickname);
	});

	$("#inputCode").on('input', function(evt) {
	  $(this).val(function (_, val) {
	  	val = val.replace(/[^a-zA-Z]/g, "");
	    return val.toUpperCase();
	  });
	});

	$("#inputNickname").on('input', function(evt) {
	  $(this).val(function (_, val) {
	  	val = val.replace(/[^a-zA-Z0-9]/g, "");
	    return val
	  });
	});

	startLoop();

});


function joinRoom(code, nickname) {
	$.ajax("http://localhost:9090/join", {
    data : JSON.stringify({ code: code, nickname: nickname }),
    contentType : 'application/json',
    type : 'POST',
    success: function(o) {
    	if(o.id != undefined && o.token != undefined && o.code != undefined) {
    		localStorage.setItem('code', o.code);
    		localStorage.setItem('id', o.id);
    		localStorage.setItem('token', o.token);
    		console.log(localStorage);
    		startLoop();
    	}
    }
	});
}

function updateRoomInfo(code) {
	$.ajax("http://localhost:9090/info", {
    data : JSON.stringify({ code: code}),
    contentType : 'application/json',
    type : 'POST',
    success: function(o) {
    	if(o.code != undefined) {
    		roomInfo = o;
    	}
    }
	});
}

function startLoop() {
	if(looping) {
		return;
	}

	looping = true;
	updateLoop();
}

function inRoom() {
	var code = localStorage.getItem('code');
	var id = localStorage.getItem('id');
	var token = localStorage.getItem('token');

	return (id != undefined && token != undefined && code != undefined)
}

function updateLoop() {
	if(!looping) {
		$('#joinform').css({"display":"default"});
		return;
	}

	looping = inRoom();
	updateRoomInfo(localStorage.getItem('code'));

	$('#joinform').css({"display":"none"});
	
	console.log("loop");
    setTimeout(updateLoop, 5000);
}