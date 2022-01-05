import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { ChatService } from '../chat.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss'],
  providers: [ChatService],
})
export class ChatComponent implements OnInit {

  private websocketSubscription : Subscription
  private message = {
    author: 'byron',
    message: 'this is a test message',
  };

  constructor(private chatService: ChatService) {
    const observer = {
      next: (x : any) => console.log('Observer got a next value: ' + x),
      error: (err : any) => console.error('Observer got an error: ' + err),
      complete: () => console.log('Observer got a complete notification'),
    };
    this.websocketSubscription = this.chatService.subscribe(observer);
  }

  ngOnInit(): void {
    console.log(this.websocketSubscription.closed);
  }

  sendMsg() {
    this.chatService.send(this.message);
  }
}
