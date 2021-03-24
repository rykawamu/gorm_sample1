const tasks = new Vue({
    el: '#task-list',
    data: {
        user: '',
        tasks: []
    },
    methods: {
        getTasks() {
            const url = 'api/tasks'
            const headers = {'Authorization': `Bearer ${this.getToken()}`}

            fetch(url, {headers}).then(response => {
                if(response.ok) {
                    return response.json()
                }
                return []
            }).then(json => {
                this.tasks = json
            })
        },
        getToken() {

            return localStorage.getItem('token')
        },
        logout() {
            localStorage.removeItem('token')
            location.href = '/'
        },
    },
    created() {
        const date = new Date()
        const claims = JSON.parse(atob(this.getToken().split('.')[1]))
        this.user = claims.name
        if(claims.exp < Math.floor(date.getTime() / 1000)) {
            this.logout()
        } else {
            this.getTasks()
        }
    },
})