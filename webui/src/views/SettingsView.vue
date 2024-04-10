<script>
// update username
export default {
    methods: {
        async setMyUsername() {
            try {
				// PUT /settings
				let username = document.getElementById('username').value;
				if (!username.match("^[a-zA-Z][a-zA-Z0-9_]{2,15}$")) {
                alert("Invalid username: 3 - 16 characters; first character must be a letter; only letters, numbers and underscores allowed");
                return;
				}
				let response = await this.$axios.put('/settings/', {username: username}, {headers: {'Authorization': `${sessionStorage.getItem('token')}`, 'Content-Type': 'application/json'}});
				let user = response.data // user_id, username
				sessionStorage.setItem('token', user.user_id);
				sessionStorage.setItem('username', user.username);
                alert(`Username correctly updated to: ${user.username}`)
            } catch (error) {
				const status = error.response.status;
        		const errorMessage = error.response.data;
        		alert(`Status (${status}): ${errorMessage}`);
            }
        }
    }
}
</script>

<template>
	<div class="background">
		<div class="container py-5">
			<div class="row justify-content-center align-items-center">
				<div class="col-md-6">
					<div class="card bg-white text-dark rounded-3">
						<div class="card-body p-5 text-center">
							<h2 class="fw-bold mb-4 text-uppercase">
								Update username
							</h2>
							<p class="text-muted">
								Please enter your new username.
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
							<div class="d-grid gap-3">
								<button
									class="btn btn-primary rounded-pill"
									type="submit"
									@click="setMyUsername"
									style="background-color: #000000">
									Update
								</button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style>
</style>
