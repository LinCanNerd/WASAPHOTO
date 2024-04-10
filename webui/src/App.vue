<!-- ./open-npm.sh -->
<!-- npm run dev -->

<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function() {
		return {
			errormsg: null,
		}
	},
	methods: {
		// login
		async doLogin() {
			try {
				// POST /session
				let username = document.getElementById('username').value;
				if (!username.match("^[a-zA-Z][a-zA-Z0-9_]{2,15}$")) {
                alert("Invalid username: 3 - 16 characters; first character must be a letter; only letters, numbers and underscores allowed");
                return;
				}
				let response = await this.$axios.post('/session/', {username: username}, {headers: {'Content-Type': 'application/json'}});
				let user = response.data // userID, username
				sessionStorage.setItem('token', user.user_id);
				sessionStorage.setItem('username', user.username);
				this.$router.replace('/home');
			} catch (error) {
				const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = `Status ${status}: ${reason}`;
			}
		}
	},
}
</script>

<template>

	<div v-if="$route.path !== '/login'">
		<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
			<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPHOTO</a>
			<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
		</header>

		<div class="container-fluid">
			<div class="row">
				<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
					<div class="position-sticky pt-3 sidebar-sticky">
						<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>General</span>
						</h6>
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink to="/home" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
									Home
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink to="/search" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
									Search
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink to="/personalProfile" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
									Profile
								</RouterLink>
							</li>
						</ul>

						<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>Secondary menu</span>
						</h6>
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink to="/settings" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg>
									Settings
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink to="/login" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
									Logout
								</RouterLink>
							</li>
						</ul>
					</div>
				</nav>

				<main class="col-md-9 ms-sm-auto col-lg-10 px-md-0">
					<RouterView />
				</main>
			</div>
		</div>
	</div>

	<div v-else>
		<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
			<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPHOTO</a>
			<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
		</header>
		<div class="background">
			<div class="container py-5">
				<div class="row justify-content-center align-items-center">
					<div class="col-md-6">
						<div class="card bg-white text-dark rounded-3">
							<div class="card-body p-5 text-center">
								<h2 class="fw-bold mb-4 text-uppercase">Welcome to wasaphoto</h2>
								<p class="text-muted">
									Please enter your username
								</p>
								<div class="form-group">
									<input
										type="text"
										id="username"
										class="form-control form-control-lg rounded-pill mb-3"
										placeholder="Username"
										required
										/>
								</div>
								<div class="d-grid gap-3" @click="doLogin">
									<button
										class="btn btn-primary rounded-pill"
										type="submit"
										style="background-color: 1">
										Login
									</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
			<ErrorMsg v-if="errormsg" :msg="errormsg" />
		</div>
	</div>

</template>

<style>
	.background {
		background-color: rgb(149, 180, 178);
		height: 100vh;
	}
	
</style>
