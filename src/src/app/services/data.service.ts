import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { APIResponse } from '../models/APIResponse'

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json'
  })
};

@Injectable({
  providedIn: 'root'
})
export class DataService {

  // Go server
  serverURL = "http://localhost:9090"
  
  constructor(
    private http : HttpClient
  ) { }

  // Get all bucket list
  getAllBuckets():Observable<APIResponse>{
    return this.http.get<APIResponse>(this.serverURL+'/getAllBuckets')
  }

  // Create new bucket
  createBucket(bucketName : string):Observable<APIResponse>{
    return this.http.post<APIResponse>(this.serverURL+'/createBucket?bucket='+bucketName,'',httpOptions)
  }

  // Delete bucket
  deleteBucket(bucketName : string):Observable<APIResponse>{
    return this.http.post<APIResponse>(this.serverURL+'/deleteBucket?bucket='+bucketName,'',httpOptions)
  }

  // Get all data
  getBucketData(bucketName : string):Observable<APIResponse>{
    return this.http.get<APIResponse>(this.serverURL+'/getAllData?bucket='+bucketName)
  }

  // Write data
  writeBucket(bucketName : string,key:any,value:any):Observable<APIResponse>{
    return this.http.post<APIResponse>(this.serverURL+'/createBucket?bucket='+bucketName+'&key='+key,value,httpOptions)
  }
}
