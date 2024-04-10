<script>
export default {
	data() {
		return {
            errormsg: null,
			photoURL: "",
			liked: false,
            likes: [],
            comments: [],
		}
	},

	props: ['pid','ownerID','username','date','likesListParent','commentsListParent','isOwner'], 

	methods: {
		getPhoto() {
			// GET /photo/{id}
			this.photoURL = __API_URL__ + `/photos/${this.pid}/`;
		},
		async deletePhoto() {
			try {
				// DELETE /photo/{id}
				await this.$axios.delete(`/photos/${this.pid}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
				this.$emit("removePhoto", this.pid);
			} catch (error) {
                const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status ' + status + ': ' + reason;
                alert(this.errormsg);
            }
		},
		visitAuthorProfile() {
            // /profiles/:username
			this.$router.push(`/profile/${this.username}`);
		},
		async likeToggle() {
			try {
				if (!this.liked) {
					// PUT /like/{id}
                    await this.$axios.put(`/likes/${this.pid}/`, null, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
					this.likes.push({user_id: sessionStorage.getItem('token'), username: sessionStorage.getItem('username')});
				} else {
					// DELETE /like/{id}
                    await this.$axios.delete(`/likes/${this.pid}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.likes = this.likes.filter(user => user.user_id != sessionStorage.getItem('token'));
				}
				this.liked = !this.liked;
			} catch (error) {
                const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status ' + status + ': ' + reason;
                alert(this.errormsg);
            }
    	},
        // on child event
		removeCommentFromList(comment_id) {
			this.comments = this.comments.filter(comment => comment.comment_id != comment_id);
		},
		addCommentToList(comment){
			this.comments.unshift(comment); // at the beginning of the list
		},
        visitLiker(username) {
            if (username != this.$route.params.username) {
                document.querySelector('.modal-backdrop').remove();
                document.querySelector('.modal').remove();
                document.body.style.overflow = 'auto';
                this.$router.push(`/profile/${username}`);
            }
        }
	},
	async mounted() {
        this.getPhoto()
        // it is a promise
        if (this.likesListParent != null) {
            this.likes = this.likesListParent
        }
        if (this.commentsListParent != null) {
            this.comments = this.commentsListParent
        }
		this.liked = this.likes.some(user => user.user_id == sessionStorage.getItem('token'));
	},
}
</script>

<template>
	<div class="container-fluid mt-3 mb-5 ">

        <UserModal
        :modalID="'likesModal'+pid" 
		:usersList="likes"
        @visitUser="visitLiker"
        />

        <CommentModal
        :modalID="'commentModal'+pid" 
		:comments="comments" 
		:isOwner="isOwner" 
		:pid="pid"
		@removeComment="removeCommentFromList"
		@addComment="addCommentToList"
		/>

        <div class="d-flex flex-row justify-content-center">
            <div class="card my-card">
                <div class="d-flex justify-content-end">
                    <button v-if="isOwner" class="my-trnsp-btn my-dlt-btn me-2" @click="deletePhoto">
						<!--trash bin-->
						<i class="fa-solid fa-trash w-100 h-100"></i>
					</button>
                </div>
                <!--photo-->
                <div class="d-flex justify-content-center photo-background-color">
                    <img :src="photoURL" class="card-img-top img-fluid">
                </div>
                <div class="card-body">
                    <div class="container">
                        <div class="d-flex flex-row justify-content-end align-items-center mb-2">
                            <!--author-->
							<button class="my-trnsp-btn m-0 p-1 me-auto" @click="visitAuthorProfile" style="background-color: bisque;">
                            	<i> From <b>{{username}}</b></i>
							</button>
                            <!--like-->
                            <button class="my-trnsp-btn m-0 p-1 d-flex justify-content-center align-items-center">
                                <i @click="likeToggle" :class="'me-1 my-heart-color w-100 h-100 fa '+(liked ? 'fa-heart' : 'fa-heart-o')"></i>
                                <i data-bs-toggle="modal" :data-bs-target="'#likesModal'+pid" class="my-comment-color ">
                                    {{likes.length}}
                                </i>
                            </button>
                            <!--comment-->
                            <button class="my-trnsp-btn m-0 p-1  d-flex justify-content-center align-items-center" 
							data-bs-toggle="modal" :data-bs-target="'#commentModal'+pid">
                                <i class="my-comment-color fa-regular fa-comment me-1"></i>
                                <i class="my-comment-color-2"> {{comments.length}}</i>
                            </button>
                        </div>
                        <div class="d-flex flex-row justify-content-start align-items-center ">
                            <p> Uploaded on <b>{{date}}</b></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.photo-background-color{
	background-color: rgb(255, 255, 255);
}
.my-card{
	width: 27rem;
	border-color: black;
	border-width: thin;
}
.my-heart-color{
	color: rgb(0, 0, 0);
}
.my-heart-color:hover{
	color: red;
}
.my-comment-color {
	color: rgb(0, 0, 0);
}
.my-comment-color:hover{
	color: rgb(255, 0, 0);
}
.my-comment-color-2{
	color:rgb(0, 0, 0)
}
.my-dlt-btn{
	font-size: 19px;
}
.my-dlt-btn:hover{
	font-size: 19px;
	color: var(--color-red-danger);
}
</style>
