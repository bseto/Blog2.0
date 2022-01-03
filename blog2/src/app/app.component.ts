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
    this.chatService.send(this.message);
  }
}
