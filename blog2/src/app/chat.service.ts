import { Injectable } from "@angular/core";
import { Observer, Subscription } from "rxjs";
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
    console.log("sending auth message: ", auth);
    this.websocketSubject.next(auth);
    //this.websocketSubject.subscribe(
      //next => console.log(next)
    //);
  }

  public subscribe(observer: Observer<any>) : Subscription {
    console.log("subscribing");
    return this.websocketSubject.subscribe(observer);
  }

  public send(data: any) {
    console.log("sending: ", data);
    this.websocketSubject.next(data);
  }

  public close() {
    console.log("closing socket");
    this.websocketSubject.complete();
  }
}