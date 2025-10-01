
<div align="center">
<h3 align="center">go-simpleDOS </h3>

  <p align="center">
    DISCLAIMER: All code shown here is solely for educational purposes, so others can better understand how a simple DOS attack with IP spoofing and SYN flooding can be implemented with Go. As it's educational, it's meant to only be ran for demonstration/testing purposes within your own local-network/devices or ones where permission has been given to you.
    <br />
    <a href="https://github.com/coherentjavi/Go-SpoofIP/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project
This project was coded with `Go 1.24.6 linux/amd64`. In practice a DOS attack launched to an arbitrary device via IP spoofing and SYN Flooding, will rarely work due to many factors. Some of these factors include most modern routers implementing RFC 2827, devices in networks sharing IPs that are later translated with NAT, extra security provided by firewalls/proxies, among others. 

To spoof the source device IP, <a href="https://pkg.go.dev/github.com/gopacket/gopacket"> GoPacket</a> and <a href= "https://pkg.go.dev/golang.org/x/net">x/net</a> (specifically ipv4) were used. 

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

Please note all devices involved need to be connected to the same network for this to run correctly. The code can be ran by going to `/src` and running all the files with `go run main.go scan.go attackPorts.go`. Root level permissions may be needed in which case `sudo go run` will work instead, if this is the case then ensure the Go environment variables are configured for root too. Alternatively the user can buid a binary locally with `go build` and execute the binary.

The TCP header and its values in `attack.go`, can be modified as needed.  


<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE 
## License

Distributed under the project_license. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>-->


