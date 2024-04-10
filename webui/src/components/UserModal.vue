<script>
export default {

	props:['modalID','usersList'],

    methods: {
        visitUser(username) {
            // /profiles/:username
            this.$emit('visitUser', username); // signal to parent
		},
    }
}
</script>

<template>
	<div class="modal fade my-modal-disp-none" :id="modalID" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-5" :id="modalID">{{ modalID.slice(0, 5) }}</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div v-for="user in usersList" :key="user.user_id">
                    <div class="modal-body">
                        <div class="container-fluid">
                            <hr>
                            <div class="row mb-2 mt-2">
                                <div class="col d-flex justify-content-center">
                                    <div class="user-mini-card card bg-transparent border-start">
                                        <div class="card-body">
                                            <h5 @click="visitUser(user.username)" class="user-mini-card-title d-flex justify-content-center ">@{{ user.username }}</h5>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
	</div>
</template>

<style>
.my-modal-disp-none{
	display: none;
}
.user-mini-card{
    color: black;
    width: 25%;
    border: none;
    border-color: black;
}
.user-mini-card-title:hover{
    cursor: pointer;
    font-weight: bold;
}
</style>
