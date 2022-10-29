# bussiness-services

### **How to install**
- Need to clone 

  - [https://github.com/vins7/user-management-services](https://github.com/vins7/user-management-services.git)
  - [https://github.com/vins7/top-up-services](https://github.com/vins7/top-up-services.git)
  - [https://github.com/vins7/emoney-services](https://github.com/vins7/emoney-services.git)

- If want to run program 
  - Run user service first then emoney, top-up and then bussiness, you can run this program with this command  


  ```
    make build
    make grpc
  ```

- If want to test program
 
  - **Open new grpc requet like this picture**
 
  ![1](https://user-images.githubusercontent.com/44151799/198817094-42778377-b2b6-4d5a-94ac-70af61821244.png)


  - **Select url port app**
  
  ![2](https://user-images.githubusercontent.com/44151799/198817974-60e8e030-e48e-4f6c-8b7e-b853018fd12e.png)
  
  
  - **Select service what you want to test, but you need to create user first**
  
  ![3](https://user-images.githubusercontent.com/44151799/198818013-173d642c-0670-42f9-8888-229300a885da.png)
  
  
  - **If you don't know what the request is, you can Generate Example Message**
  
  ![5](https://user-images.githubusercontent.com/44151799/198818082-61492266-65dc-49f6-bcca-fd582dd7fce3.png)
  
  
  - **And then the request will appear**
  
  ![4](https://user-images.githubusercontent.com/44151799/198818156-8caac048-67f1-4279-a092-5670bd8f6f67.png)
