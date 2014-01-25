/*
 * @author sheppard(ysf1027@gmail.com) 2014-01-02
 */

var socket;
var msgs = [];

var msgHandles = {};
msgHandles['HCChat'] = function(msg) {
  addToPanel(msg.Name, msg.Content);
  refreshPanel();
};
msgHandles['HCRoomList'] = function(msg) {
  var rooms = $('#panel-roomlist');
  /*
  $.each(msg.Rooms, function(index, room) {
    rooms.append('<a href="' + room.Href + '">' + room.Name + '</a>' + '  <span class="badge">' + room.OnlineCount + '</span>');
    rooms.append('<br>');
  });
  */
};

$(document).ready(function() {
  connect();

  var btn_chat = $('#btn-chat');
  $.each(btn_chat.children(), function(_, item) {
    var item = $(item);
    item.click(function() {
      send({Type: 'chat', Data: $(this).text()});
    });
  });

  $('#btn-pvpRoom').click(function() {
    send({Type: 'enterRoom', Data: 'pvp'});
  });
  $('#btn-pveRoom').click(function() {
    send({Type: 'enterRoom', Data: 'pve'});
  });
  $('#btn-exitRoom').click(function() {
    send({Type: 'exitRoom', Data: ''});
  });
  $('#btn-start').click(function() {
    send({Type: 'start', Data: ''});
  });
  $('#btn-giveUp').click(function() {
    send({Type: 'giveUp', Data: ''});
  });
});

function connect() {
  var userName = $('#userName').text();
  socket = new WebSocket('ws://'+window.location.host+'/hall/hall/socket?user=' + userName);

  socket.onmessage = function(event) {
    console.log(event);
    var pack = JSON.parse(event.data);
    var handle = msgHandles[pack.Type];
    if(handle) {
      handle(pack.Data);
    }
  };

  socket.onclose = function() {
    console.log('disconnected');
  };
};

function addToPanel(prefix, msg) {
  msgs.push({prefix: prefix, msg: msg});
};

function refreshPanel() {
  var panel = $('#panel-chat');
  panel.empty();
  $.each(msgs, function(index, item) {
    panel.append('<strong>' + item.prefix + ':</strong> ');
    panel.append(item.msg).append('<br>');
  });
};

function refreshRoomList() {
  var rooms = $('#list-room');
}

function send(msg) {
  socket.send(JSON.stringify(msg));
}
