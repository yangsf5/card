/*
 * @author sheppard(ysf1026@gmail.com) 2014-01-02
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
  rooms.empty();
  _.each(msg.Rooms, function(room) {
    rooms.append('<a href="' + room.Href + '" target="_blank">' + room.Name + '</a>' + '  <span class="badge">' + room.OnlineCount + '</span>');
    rooms.append('<br>');
  });
};

$(document).ready(function() {
  connect();

  var btn_test = $('#btn-test');
  btn_test.click(function() {
    var content = btn_test.text(); 
    socket.send(content);
    return false;
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
  _.each(msgs, function(item) {
    panel.append('<strong>' + item.prefix + ':</strong> ');
    panel.append(item.msg).append('<br>');
  });
};

function refreshRoomList() {
  var rooms = $('#list-room');
}
