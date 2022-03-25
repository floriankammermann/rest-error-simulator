import { HttpHeaders } from '@angular/common/http';
import { Component, ElementRef, ViewChild } from '@angular/core';
import { HttpserviceService } from './services/httpservice.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'rest-error-simulator-frontend';
  error_response = "";

  control_response = {
    "responseCodeSuccess": Number,
    "responseCodeFailure": Number,
    "responseCodeSuccessFailureRatio": Number,
    "ratioModulo": Number,
    "requestCounter": Number
}

  @ViewChild('errorratio', { read: ElementRef }) errorRatio:
  | ElementRef
  | undefined;

  @ViewChild('latency', { read: ElementRef }) latency:
  | ElementRef
  | undefined;

  constructor(private backendData: HttpserviceService){}

  onSubmitErrorRatio() {
    let _errorRatio = this.errorRatio?.nativeElement.value;
    if (_errorRatio > 0 && _errorRatio <= 100) {
      this.error_response = "";
      this.backendData.sendErrorRatio(_errorRatio).subscribe((response: any) => {
      });
    } else {
      this.error_response = "Error Ratio muss zwischen 1 und 100 sein."
    }
  }

  onSubmitLatency() {
    let _latency = this.latency?.nativeElement.value;
    if (_latency > 0 && _latency <= 10000) {
      this.error_response = "";
      this.backendData.sendLatency(_latency).subscribe((response: any) => {
      })
    } else {
      this.error_response = "Latency muss zwischen 1 und 10000 sein."
    }
  }

  getControls() {
    this.backendData.getControls().subscribe((response: any) => {
      console.log(response);

    })
  }
}
