<script>
// search user by username
export default {
	data: function() {
		return {
			users: [],
		}
	},
	methods: {
		async searchUserByUsername() {
			try {
				// GET /users/?username=
                let username = document.getElementById('username').value;
                let response = await this.$axios.get('/search/', {params: {username: username}, headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.users = response.data
			} catch (error) {
                const status = error.response.status;
        		const reason = error.response.data;
                alert(`Status` + status + `:` + reason);
			}
		},
        visitProfile(username) {
            // /profiles/:username
			this.$router.push(`/profile/${username}`);
		},
	},
}
</script>

<template>
    <div class="search-container">
        <form @submit.prevent="searchUserByUsername">
            <input id="username" class="search-box" type="text" placeholder="Username">
            <button class="search-button" type="submit">Search</button>
        </form>
    </div>

    <div v-for="user in users" :key="user.user_id">
        <div class="modal-body">
            <div class="container-fluid">
                <div class="row mb-2 mt-2">
                    <div class="col d-flex justify-content-center">
                        <div class="user-mini-card card bg-transparent border-start">
                            <div class="card-body">
                                <h5 @click="visitProfile(user.username)" class="user-mini-card-title d-flex justify-content-center ">@{{ user.username }}</h5>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.search-container {
    text-align: center;
    padding-top: 13px;
}
.search-box {
    padding: 5px;
    border: 1px solid #ccc;
    border-radius: 5px;
    width: 300px;
    max-width: 80%;
}
.search-button {
    padding: 6px 10px;
    background-color: #4489ce;
    color: #fff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}
</style>
