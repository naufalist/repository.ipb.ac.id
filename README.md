<div id="top"></div>
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/naufalist/repository.ipb.ac.id">
    <img src="images/Logo-IPB-University-Vertical.png" alt="Logo" height="125">
  </a>

  <h3 align="center">IPB Repository Scrapper & Downloader</h3>

  <p align="center">
    Simple repository scrapper & downloader written in Go
    <br />
<!--     <a href="https://github.com/naufalist/repository.ipb.ac.id"><strong>Explore the docs »</strong></a>
    <br />
    <br /> -->
    <a href="https://tools.naufalist.com/ipb-repository-downloader">View Demo</a>
    ·
    <a href="https://github.com/naufalist/repository.ipb.ac.id/issues">Report Bug</a>
    ·
    <a href="https://github.com/naufalist/repository.ipb.ac.id/issues">Request Feature</a>
  </p>
</div>


## Note (Important!)
As of this date (07/04/2022), this application is no longer allowed to run.
However, you could be able to run it locally.
The source code is still open-source, until someone asked me to remove it 😂
I'm done with this. Thank you.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- ABOUT THE PROJECT -->
## About
<p align="center">
  <img src="https://github.com/naufalist/repository.ipb.ac.id/blob/main/images/screenshot.png?raw=true" alt="Screenshot"/>
</p>

In the internet, there are many website that provide research journals and/or theses, such as [IPB Repository](https://repository.ipb.ac.id/), [LIPI](http://isjd.pdii.lipi.go.id/), [e-Journal Perpusnas](https://ejournal.perpusnas.go.id/), and [Google Scholar](https://scholar.google.com/). IPB repository is one of the most popular university repositories on the internet. Unfortunately, only people with access can download locked files (students, lecturers, or staff). Therefore, I create this app to help anyone who wants to get free access to it. (***my credentials are stored in the server-side***).


### Built With

* [Go](https://golang.org/) `v1.16`
* [Bootstrap](https://getbootstrap.com) `v5`
* [goquery](https://github.com/PuerkitoBio/goquery) `v1.7.1`
* [godotenv](https://github.com/joho/godotenv) `v1.4.0`

### Prerequisites

To run this app, you need Go in your environment.
Please go to the following link: [https://golang.org/](https://golang.org/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Installation

1. Clone the repository.
   ```sh
   git clone https://github.com/naufalist/repository.ipb.ac.id.git
   ```
2. Copy **.env.example** to **.env**.
   ```sh
   cp .env.example .env
   ```
3. In **.env**, you should change the credentials.
   ```sh
   LDAP_USERNAME=CHANGE_WITH_YOUR_LDAP_USERNAME
   LDAP_PASSWORD=CHANGE_WITH_YOUR_LDAP_PASSWORD
   ```
4. Run the app
   ```go
   go run main.go
   ```
5. Now, this app can be accessed in http://localhost:9000.

6. **(Optional)**. If you want to use a different port, you can change this value in `.env` file.
   ```sh
   PORT=9000
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Access this site: http://localhost:9000.
1. Copy-paste the repository url into the form. Example repository url: https://repository.ipb.ac.id/handle/123456789/39884.
2. Click the generate button.
3. If success, information from the repository will be displayed below the form.
4. To download repository file, you can click each download button.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
<!-- ## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>
 -->


<!-- CONTACT -->
## Contact

[@naufalist](https://twitter.com/naufalist) - contact.naufalist@gmail.com

<p align="right">(<a href="#top">back to top</a>)</p>
