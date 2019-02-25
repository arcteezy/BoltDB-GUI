import { Component, OnInit } from "@angular/core";
import { DataService } from "../../services/data.service";
import { NgForm } from "../../../../node_modules/@angular/forms";

@Component({
  selector: "app-buckets",
  templateUrl: "./buckets.component.html",
  styleUrls: ["./buckets.component.css"]
})
export class BucketsComponent implements OnInit {
  // Bucket list
  buckets: string[];

  constructor(private dataService: DataService) {}

  ngOnInit() {
    // Get all bucket list
    this.getAllBuckets();
  }

  // Get all buckets
  getAllBuckets() {
    this.dataService.getAllBuckets().subscribe(
      response => {
        console.log("Get all buckets response :", response);
        if (response != null) {
          this.buckets = JSON.parse(response.body);
        }
      },
      error => {
        console.log("Error :", error.error);
      }
    );
  }

  // Create bucket
  createBucket(form: NgForm) {
    console.log("Create new bucket :", form.value);
    if (form.value.bucketName == "") {
      alert("Bucket name not entered");
    } else {
      // Service call
      this.dataService.createBucket(form.value.bucketName).subscribe(
        response => {
          console.log("Create bucket response :", response);
          if (response != null) {
            if (response.body == "SUCCESSFUL") {
              this.getAllBuckets();
            }
          }
        },
        error => {
          console.log("Error :", error.error);
        }
      );
    }
  }

  // Delete bucket
  deleteBucket(bucket: string) {
    console.log("Delete bucket :", bucket);
    // Service call
    this.dataService.deleteBucket(bucket).subscribe(
      response => {
        console.log("Delete bucket response :", response);
        if (response != null) {
          if (response.body == "SUCCESSFUL") {
            this.getAllBuckets();
          }
        }
      },
      error => {
        console.log("Error :", error.error);
      }
    );
  }
}
