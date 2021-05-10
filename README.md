# Customer-Billing-System
This repository contains the code of customer billing system in which a buyer can buy different things like vegetables and fruits and see the shopping record by entering the name or the phone number from the file and get access to the account details.

This software is completely written in Go language and it contains five functions that will perform different tasks and the main function that will call the functions. All the functions are managed in two parts. The first part is **start shopping** in which a user will buy a product and the record will be saved in a **CSV** file. The second part is **Check the shopping record** in which a user will access the account by simply typing the name or the phone number.

---

### Table of Contents

The headers will be used to reference the location of destination.

- [The First Part](#the-first-part)
- [The Second Part](#the-second-part)
- [Author Info](#author-info)

---

# The First Part
The first part as I have written above will show the shopping option. This part contains four functions. 
1. ```listItems()``` will show all the product names and their prices so that a user can easily choose from them.
2. ```oneLine()``` will accept the input with spaces. It means if a user entered the firstname and the lastname. Then the space should not be a barrier and both names with spaces will be accepted. Although this function is not directly used in ```main()``` function but it will perform the task in ```addCustomer()```.
3. ```addCustomer()``` will take the details from the customer like his name, address, phone number, product which he wants to buy, and the payment. After entering the details, the item will be given to him and the price will be deducted.
4. ```writeFile()``` will write the all the data in a **CSV** file as a record.

[Back To The Top](#Customer-Billing-System)

---

# The Second Part
The second part will check the shopping record. It contains only one function.

1. ```showData()``` will take the name or phone number of a user and the data from the **CSV** will be shown to him if it is present.

[Back To The Top](#Customer-Billing-System)

---

## Author Info

- YouTube - [ibilalkayy](https://www.youtube.com/channel/UCBLTfRg0Rgm4FtXkvql7DRQ)
- LinkedIn - [@ibilalkayy](https://www.linkedin.com/in/ibilalkayy/)

[Back To The Top](#Customer-Billing-System)
