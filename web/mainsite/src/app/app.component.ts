import { HttpClient } from '@angular/common/http';
import { AfterViewInit, Component } from '@angular/core';
import { BookInfo, BookstoreService } from '../apis';



@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements AfterViewInit {
  public title = 'mainsite';


  public customerNameInputValue = ""
  public bookList: BookInfo[] = []
  public userInventory: BookInfo[] = []

  public ngAfterViewInit() {
    console.log("Hello")
  }

  constructor(private bookStore: BookstoreService){ }


  public async onFetchDataButtonClick() {
    const result = await this.bookStore.bookstoreGetBookInfoList().toPromise()
    this.bookList = result.bookList

    const userInventoryResult = await this.bookStore.bookstoreMyInventory(this.customerNameInputValue).toPromise()
    this.userInventory = userInventoryResult.bookList
  }

  public async onAddToCartButtonClick(id: string) {

    if (this.customerNameInputValue === "") {
      alert("ไม่ได้ ต้องใส่ customer name นะ")
      return
    }


    const isConfirm = confirm("ยืนยัน การสั่งซื้อ")
    if (!isConfirm) {
      return
    }

    await this.bookStore.bookstorePurchaseBook({
      bookId:id,
      customerName: this.customerNameInputValue,
      numberOfOrder: 1,
    }).toPromise()
  }
}
