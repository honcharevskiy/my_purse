const { createApp } = Vue


async function fetchData(){
  let response = await fetch("http://localhost:8080/monobank/personal");
  let data = await response.json();
  data = JSON.stringify(data);
  data = JSON.parse(data);
  return data;
 }


async function main() {
  let abc = await fetchData(); // here the data will be return.
  console.log(abc); // you are using async await then no need of .then().

  createApp({
    data() {  
      return {
        message: abc
      }
    }
  }).mount('#app')
}


main()