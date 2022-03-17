import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpserviceService {

  constructor(private httpClient: HttpClient) { }

  sendErrorRatio(errorRatio: Number): any {
    return this.httpClient
      .post<any>(
        "http://localhost:8080/control/error?errorratio=" + errorRatio,
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
        "http://localhost:8080/control/latency?latencyinms=" + latency,
        latency
      )
      .pipe(
        catchError((error) => {
          return throwError('Error Ratio API not working')
        })
      );
  }
}
