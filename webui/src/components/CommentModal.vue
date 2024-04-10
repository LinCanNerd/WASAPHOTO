<script>
export default {	
	data(){
		return{
            token: sessionStorage.getItem('token'),
			text:"",
		}
	},

	props:['modalID','comments','isOwner','pid'],

	methods: {

		async commentPhoto() {
            try {
                // POST /photo/{pid}/comment
                let response = await this.$axios.post(`/photos/${this.pid}/comment/`, this.text, {headers: {'Authorization': `${sessionStorage.getItem('token')}`, 'Content-Type': 'text/plain'}});
                let comment = response.data;
                this.$emit('addComment', comment); // signal to parent
                console.log("new comment:" ,this.comments);
                this.text = "";
            } catch(error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status` + status + `: ` + reason;
                alert(this.errormsg);
            }
        },
        async uncommentPhoto(comment_id) {
            try {
                // DELETE /photo/{pid}/comment/{comment_id}
                await this.$axios.delete(`/photos/${this.pid}/comment/${comment_id}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.$emit('removeComment', comment_id); // signal to parent
            } catch(error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status` + status + `: ` + reason;
                alert(this.errormsg);
            }
        },

	},
    mounted() {
        // Log comments when the component is mounted
        console.log("Comments List:", this.comments);
    }
}
</script>

<template>
    <div class="modal fade my-modal-disp-none" :id="modalID" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-5" :id="modalID">Comments</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-footer d-flex justify-content-center w-100">
                    <div class="row w-100 ">
                        <div class="col-10">
                            <div class="mb-3 me-auto">
                                <textarea class="form-control" id="exampleFormControlTextarea1" 
								placeholder="Add a comment..." rows="1" maxLength="2200" v-model="text"></textarea>
                            </div>
                        </div>
                        <div class="col-2 d-flex align-items-center">
                            <button type="button" class="btn btn-primary" 
							@click.prevent="commentPhoto" 
							:disabled="text.length < 1 || text.length > 2200">
							Send
							</button>
                        </div>
                    </div>
                </div>
                <div class="modal-body" style="overflow-y: auto;">
                    <div v-for="comment in comments" :key="comment.comment_id" class="comment-container">
                        <div class="container-fluid">
                            <div class="row">
                                <div class="col-10">
                                    <h5>@<b>{{comment.username}}</b></h5>
                                </div>
                                <div class="col-2">
                                    <button v-if="token == comment.user_id || isOwner" class="btn my-btn-comm" @click="uncommentPhoto(comment.comment_id)">
                                        <i class="fa-regular fa-trash-can my-trash-icon"></i>
                                    </button>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-12">
                                   <h4>{{comment.text}}</h4> 
                                   {{ comment.date }}
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
.my-btn-comm{
    border: none;
}
.my-btn-comm:hover{
    border: none;
    color: red;
    transform: scale(1.1);
}

.scrollable-content {
    max-height: 300px; /* Adjust the height as needed */
    overflow-y: auto;
}

.comment-container {
    margin-bottom: 30px;
    border-radius: 1cm;
    border-color: black; /* Adjust the value to increase or decrease the space between comments */
}
</style>
