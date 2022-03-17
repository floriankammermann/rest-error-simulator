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
        console.log(response);
      });
    } else {
      this.error_response = "Error Ratio ist unter 0 oder über 100. Wähle eine Zahl zwishen 0 und 100."
    }
  }
  
  onSubmitLatency() {
    let _latency = this.latency?.nativeElement.value;
    if (_latency > 0 && _latency <= 10000) {
      this.error_response = "";
      this.backendData.sendLatency(_latency).subscribe((response: any) => {
        console.log(response);
      })
    } else {
      this.error_response = "Latency ist unter 0. Wähle eine Zahl über 0."
    }
  }
}
