import { Injectable } from '@angular/core';
import { Observable, Subject, Observer } from 'rxjs';


@Injectable()
export class WebsocketService {
  constructor() {}

  private subject = new Subject<MessageEvent>();

  public connect(url : any): Subject<MessageEvent> {
    if (!this.subject) {
      this.subject = this.create(url);
      console.log("Successfully connected: " + url);
    }
    return this.subject;
  }

  private create(url : any): Subject<MessageEvent> {
    let ws = new WebSocket(url);

    let observable = new Observable((obs: Observer<MessageEvent>) => {
      ws.onmessage = obs.next.bind(obs);
      ws.onerror = obs.error.bind(obs);
      ws.onclose = obs.complete.bind(obs);
      return ws.close.bind(ws);
    });
    observable.subscribe( {
      next: (data: Object) => {
        if (ws.readyState === WebSocket.OPEN) {
          ws.send(JSON.stringify(data));
        }
      }
    });
    let subject = new Subject<MessageEvent>();
    observable.subscribe(subject);
    return subject;

    //let observer = {;
      //next: (data: Object) => {
        //if (ws.readyState === WebSocket.OPEN) {
          //ws.send(JSON.stringify(data));
        //}
      //}
    //};
    //return Subject.create(observer, observable);
  }
}