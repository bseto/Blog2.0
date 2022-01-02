import { Injectable } from "@angular/core";
import { Subject } from "rxjs";
import { map } from "rxjs/operators";
import { WebSocketSubject } from 'rxjs/webSocket';

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
  private socket: WebSocket;

  public constructor() {
    this.socket = new WebSocket("ws://localhost:8081/ws/1");
    this.socket.onopen = event => {
      let auth = new Auth(true, "");
      console.log("sending socket the empty auth");
      this.socket.send(JSON.stringify(auth));
    };
    this.socket.onclose = event => {
      console.log("on close: ", event);
    };
    this.socket.onmessage = event => {
      console.log("on close: ", JSON.parse(event.data));
    };
  }

  public send(data: string) {
    console.log("sending from somewhere idk where: ", data);
    this.socket.send(data);
  }

  public close() {
    this.socket.close();
  }
}