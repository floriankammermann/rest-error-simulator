import { HttpHeaders } from '@angular/common/http';
import { Component, ElementRef, HostListener, ViewChild } from '@angular/core';
import {v4 as uuidv4} from 'uuid';
import { HttpserviceService } from './services/httpservice.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'rest-error-simulator-frontend';
  error_response = "";
  myuuid = "";

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


  @HostListener("window:beforeunload", ["$event"]) unloadHandler(event: Event) {
    this.releaseUUID()
}

  onSubmitErrorRatio() {
    let _errorRatio = this.errorRatio?.nativeElement.value;
    if (_errorRatio > 0 && _errorRatio <= 100) {
      this.error_response = "";
      if (this.myuuid == "") {
        this.error_response = "Du musst die Instance zuerst claimen"
      } else {
        this.backendData.sendErrorRatio(_errorRatio, this.myuuid).subscribe((_: any) => {
        });
      }
    } else {
      this.error_response = "Error Ratio muss zwischen 1 und 100 sein."
    }
  }

  onSubmitLatency() {
    let _latency = this.latency?.nativeElement.value;
    if (_latency > 0 && _latency <= 10000) {
      this.error_response = "";
      if (this.myuuid == "") {
        this.error_response = "Du musst die Instance zuerst claimen"
      } else {
        this.backendData.sendLatency(_latency, this.myuuid).subscribe((_: any) => {
        }, (_: any) =>{
          this.myuuid = ""
          this.error_response = "UUID ist abgelaufen. Versuchen Sie erneut die Instance zu claimen."
        })
      }
    } else {
      this.error_response = "Latency muss zwischen 1 und 10000 sein."
    }
  }

  getControls() {
    this.backendData.getControls().subscribe((response: any) => {
    })
  }

  getUUID() {
    let tempUUID = uuidv4()
    this.backendData.sendUUID(tempUUID).subscribe((response: any) => {
      if(response != "Instance is already claimed") {
        this.myuuid = response
        this.error_response = ""
      } else  {
        this.error_response = response
      }
    })
  }

  releaseUUID() {
    this.backendData.removeUUID(this.myuuid).subscribe((response: any) => {
      this.myuuid = ""
    })
  }

  ngOnDestroy(): void {
    this.releaseUUID()
  }
}
