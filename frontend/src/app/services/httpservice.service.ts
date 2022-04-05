import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, Observable, throwError } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class HttpserviceService {

  headers: HttpHeaders = new HttpHeaders()
  .append('Content-Type', 'application/json')


  constructor(private httpClient: HttpClient) { }

  sendErrorRatio(errorRatio: Number, uuid: string): any {
    return this.httpClient
      .post<any>(
        environment.backendUrl + "/control/error?errorratio=" + errorRatio + "&clientUUID=" + uuid,
        errorRatio
      )
      .pipe(
        catchError((error) => {
          return throwError('Error Ratio API not working')
        })
      );
  }

  sendLatency(latency: Number, uuid: string): any {
    return this.httpClient
      .post<any>(
        environment.backendUrl + "/control/latency?latencyinms=" + latency + "&clientUUID=" + uuid,
        latency
      )
      .pipe(
        catchError((error) => {
          return throwError('Latency API not working')
        })
      );
  }

  sendUUID(uuid: string): any {
    return this.httpClient
      .get<any>(
        environment.backendUrl + "/control/uuid?clientUUID=" + uuid,
      )
  }

  removeUUID(uuid: string): any {
    return this.httpClient
      .get<any>(
        environment.backendUrl + "/control/uuid/delete?clientUUID=" + uuid,
      )
  }

  getControls(): any {
    return this.httpClient.get<any>( environment.backendUrl + "/control", {
      headers: this.headers
    }).pipe(
      catchError((error) => {
        return throwError("GET not working")
      })
      )
  }
}
