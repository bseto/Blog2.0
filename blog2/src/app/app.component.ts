import { Component } from '@angular/core';
import { ChatService } from "./chat.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers:[ChatService]
})
export class AppComponent {
  constructor(private chatService: ChatService) {}

  private message = {
    author: "byron",
    message: "this is a test message"
  };

  sendMsg() {
    console.log("new message from client to websocket: ", this.message);
    this.chatService.send(JSON.stringify(this.message));
  }
}
