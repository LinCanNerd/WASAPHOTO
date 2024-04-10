<script>
export default {
	methods: {
        // redirects to /profiles/:username
		async redirect() {
		    this.$router.replace(`/profile/${sessionStorage.getItem('username')}`);
		},
	},
	mounted() {
		this.redirect()
	}
}
</script>

<template>
	<!--child elem required-->
	<span></span>
</template>

<style>
</style>
