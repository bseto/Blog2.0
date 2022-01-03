import { Injectable } from "@angular/core";
import { Subject } from "rxjs";
import { map } from "rxjs/operators";
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

const CHAT_URL = "ws://localhost:8081/ws/1";

export interface Message {
  author: string;
  message: string;
}

export class Auth {
  ContainsToken: boolean;
  Token: string;
  public constructor(containsToken: boolean, Token: string) {
    this.ContainsToken = containsToken;
    this.Token = Token;
  }
}

@Injectable()
export class ChatService {

  private websocketSubject: WebSocketSubject<any>

  public constructor() {
    this.websocketSubject = webSocket("ws://localhost:8081/ws/1");
    let auth = new Auth(false, "");
    let authMessage = JSON.stringify(auth);
    console.log("sending auth message: ", auth);
    this.websocketSubject.next(auth);
    this.websocketSubject.subscribe(
      msg => console.log("message received: " + msg),
      err => console.log("error: " + err),
      () => console.log("connection closed")
    );
  }

  public send(data: string) {
    this.websocketSubject.next(data);
  }

  public close() {
    this.websocketSubject.complete();
  }
}