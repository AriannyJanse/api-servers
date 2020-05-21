import axios from 'axios'

const url = 'http://localhost:8083/servers'

function getServers() {
    return  axios
        .get(url)
        .then(response => response.data.items)
        .catch((error) => {
            console.log("this error: "+error);
        })
}

function getServerByDomain(domain) {
    return  axios
        .get(`${url}/${domain}`)
        .then(response => response.data)
        .catch((error) => {
            console.log("this error: "+error);
        })
}

export default {
    getServers,
    getServerByDomain,
}