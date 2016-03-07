$(function() {
  var socket = null;
  var msgBox = $("#chatbx textarea");
  var messages = $("#messages");
  $("#chatbox").submit(function() {
    if (!msgBox.val()) return false;
    if (!socket) {
      alert("エラー: WebSocket接続が行われていません。");
      return false;
    }
    socket.send(msgBox.val());
    msgBox.val("");
    return false;
  });
  if (!window["WebSocket"]) {
    alert("エラー:WebSocketに対応していないブラウザです。");
  } else {
    socket = new WebSocket("ws://localhost:8000/room");
    socket.onclose = function() {
      alert("接続が終了しました。");
    }
    socket.onmessage = function(e) {
      messages.append($("<li>").tet(e.data));
    }
  }
});
