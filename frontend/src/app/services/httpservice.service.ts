import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, throwError } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class HttpserviceService {

  headers: HttpHeaders = new HttpHeaders()
  .append('Content-Type', 'application/json')


  constructor(private httpClient: HttpClient) { }

  sendErrorRatio(errorRatio: Number): any {
    return this.httpClient
      .post<any>(
        environment.backendUrl + "/control/error?errorratio=" + errorRatio,
        errorRatio
      )
      .pipe(
        catchError((error) => {
          return throwError('Error Ratio API not working')
        })
      );
  }

  sendLatency(latency: Number): any {
    return this.httpClient
      .post<any>(
        environment.backendUrl + "/control/latency?latencyinms=" + latency,
        latency
      )
      .pipe(
        catchError((error) => {
          return throwError('Error Ratio API not working')
        })
      );
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
