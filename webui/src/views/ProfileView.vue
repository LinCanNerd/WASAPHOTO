<script>
// get user profile
export default {
	data: function() {
		return {
            errormsg: null,

            // getUserProfile
			username: "",
            photosCount: 0,
            followersCount: 0,
            followingCount: 0,
            isOwner: false,
            doIFollowUser: false,
            isInMyBannedList: false,
            amIBanned: false,

            // getPhotosList
            photosList: [],

            // getFollowersList
            followerList: [],

            // getFollowingsList
            followingList: [],

            userExists: false,
            user_id: 0,
		}
	},
    watch: {
        // property to watch
        pathUsername(newUName, oldUName) {
            if (newUName !== oldUName){
                this.getUserProfile()
            }
        }
    },
    computed: {
        pathUsername() {
            return this.$route.params.username
        },
    },
    methods: {
        async getUserProfile() {
            if (this.$route.params.username === undefined) {
                return
            }
            try {
                let username = this.$route.params.username;
                let response = await this.$axios.get(`/search/`,{params: {username: username}, headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.user_id = response.data[0].user_id;
                response = await this.$axios.get(`/users/${this.user_id}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                let profile = response.data;
                this.username = profile.username;
                if (profile.photoList != null) {
                    this.photosCount = profile.photoList.length;
                }else{
                    this.photosCount = 0;
                }

                if (profile.followerList != null) {
                    this.followersCount = profile.followerList.length;
                    for (let i = 0; i < profile.followerList.length; i++) {
                        if (profile.followerList[i].user_id === sessionStorage.getItem('token')) {
                            this.doIFollowUser = true;
                            break;
                        }
                    }
                }else{
                    this.followersCount = 0;
                }
                if (profile.followingList != null) {
                    this.followingCount = profile.followingList.length;
                }else{
                    this.followingCount = 0;
                }
                if (profile.user_id === sessionStorage.getItem('token')) {
                    this.isOwner = true;
                }
                
                let hostresponse = await this.$axios.get(`/users/${sessionStorage.getItem('token')}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                let hostprofile = hostresponse.data;
                
                if(hostprofile.bannedList != null) {
                    for (let i = 0; i < hostprofile.bannedList.length; i++) {
                        if (hostprofile.bannedList[i].user_id === this.user_id) {
                            this.isInMyBannedList = true;
                            break;
                        }
                    }
                }

                if(profile.bannedList != null) {
                    for (let i = 0; i < profile.bannedList.length; i++) {
                        if (profile.bannedList[i].user_id === sessionStorage.getItem('token')) {
                            this.amIBanned = true;
                            break;
                        }
                    }
                }
                
            
                this.userExists = true;
                if (!this.isInMyBannedList && !this.amIBanned) {
                    await this.getPhotosList();
                    this.getFollowersList();
                    this.getFollowingsList();
                }
            } catch (error) {
                const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
		async followBtn() {
            try {
                if (this.doIFollowUser) { 
                     // DELETE /following/{uid}
                    await this.$axios.delete(`/following/${this.user_id}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.getUserProfile();
                } else {
                    // PUT /following/{uid}
                    await this.$axios.put(`/following/${this.user_id}/`, null, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.getUserProfile();
                }
                this.doIFollowUser = !this.doIFollowUser
            } catch (error) {
                const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
		},
        async banBtn() {
            try {
                if (this.isInMyBannedList) {
                    // DELETE /banned/{uid}
                    await this.$axios.delete(`/ban/${this.user_id}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.isInMyBannedList = false;
                    this.getUserProfile();
                } else {
                    // PUT /banned/{uid}
                    await this.$axios.put(`/ban/${this.user_id}/`,null, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                    this.getUserProfile();
                }
            } catch (error) {
                const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
		},
        async uploadPhoto() {
            try {
                let file = document.getElementById('fileUploader').files[0];
                const reader = new FileReader();
                reader.readAsArrayBuffer(file); // stored in result attribute
                reader.onload = async () => {
                    // POST /photos/
                    let response = await this.$axios.post('/photos/', reader.result, {headers: {'Authorization': `${sessionStorage.getItem('token')}`, 'Content-Type': 'image/*'}});
                    this.photosList.unshift(response.data); // at the beginning of the list
                    this.photosCount += 1;
                }
            } catch (error) {
                const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
        async getPhotosList() {
            try {
                // GET /users/{uid}/photos/
                let response = await this.$axios.get(`/users/${this.user_id}/photo/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.photosList = response.data === null ? [] : response.data;
            } catch (error) {
				const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
        async getFollowersList() {
            try {
                // GET /users/{uid}/followers/
                let response = await this.$axios.get(`/users/${this.user_id}/followers/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.followerList = response.data;
                if (this.followerList===null){
                    this.followerList = [];
                }
            } catch (error) {
				const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
        async getFollowingsList() {
            try {
                // GET /users/{uid}/followings/
                let response = await this.$axios.get(`/users/${this.user_id}/following/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.followingList = response.data;
                if (this.followingList===null){
                    this.followingList = [];
                }
            } catch (error) {
				const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = 'Status' + status + ': ' + reason;
            }
        },
        // on child event
        removePhotoFromList(pid){
			this.photosList = this.photosList.filter(photo => photo.pid != pid);
            this.photosCount -= 1;
		},
        visitUser(username) {
            if (username != this.$route.params.username) {
                this.$router.push(`/profile/${username}`);
            }
        }
    },
    mounted() {
        this.getUserProfile();
    }
}
</script>

<template>

    <UserModal
    :modalID="'usersModalFollowers'" 
    :usersList="followerList"
    @visitUser="visitUser"
    />

    <UserModal
    :modalID="'usersModalFollowing'" 
    :usersList="followingList"
    @visitUser="visitUser"
    />

    <div class="container-fluid" v-if="userExists && !amIBanned">
        <div class="row">
            <div class="col-12 d-flex justify-content-center">
                <div class="card w-50 container-fluid">
                    <div class="row">
                        <div class="col">
                            <div class="card-body d-flex justify-content-between align-items-center">
                                <h5 class="card-title p-0 me-auto mt-auto">@{{username}}</h5>

                                <button v-if="!isOwner && !isInMyBannedList" @click="followBtn" class="btn btn-success ms-2">
                                    {{doIFollowUser ? "Unfollow" : "Follow"}}
                                </button>

                                <button v-if="!isOwner" @click="banBtn" class="btn btn-danger ms-2">
                                    {{isInMyBannedList ? "Unban" : "Ban"}}
                                </button>
                            </div>
                        </div>
                    </div>

                    <div v-if="!isInMyBannedList" class="row mt-1 mb-1">
                        <button class="col-4 d-flex justify-content-center btn-foll">
                            <h6 class="ms-3 p-0 ">Posts: {{photosCount}}</h6>
                        </button>
                    
                        <button class="col-4 d-flex justify-content-center btn-foll">
                            <h6 data-bs-toggle="modal" :data-bs-target="'#usersModalFollowers'">
                                Followers: {{followersCount}}
                            </h6>
                        </button>
                    
                        <button class="col-4 d-flex justify-content-center btn-foll">
                            <h6 data-bs-toggle="modal" :data-bs-target="'#usersModalFollowing'">
                                Following: {{followingCount}}
                            </h6>
                        </button>
                    </div>
                </div>
            </div>
        </div>


        <div class="row">
            <div class="container-fluid mt-3">
                <div class="row ">
                    <div class="col-12 d-flex justify-content-center">
                        <h2>Posts</h2>
                        <input id="fileUploader" type="file" class="profile-file-upload" @change="uploadPhoto" accept=".jpg, .png">
                        <label v-if="isOwner" class="btn my-btn-add-photo ms-2 d-flex align-items-center" for="fileUploader"> Add </label>
                    </div>
                </div>
                <div class="row ">
                    <div class="col-3"></div>
                    <div class="col-6">
                        <hr class="border border-dark">
                    </div>
                    <div class="col-3"></div>
                </div>
            </div>
        </div>

        <div class="row">
            <div class="col">
                <div v-if="!isInMyBannedList && photosCount>0">
                    <Photo v-for="photo in photosList"
                    :key="photo.pid"
                    :pid="photo.pid"
                    :ownerID="photo.user_id"
                    :username="photo.username"
                    :date="photo.date"
                    :likesListParent="photo.likes"
                    :commentsListParent="photo.comments"
                    :isOwner="isOwner"
                    @removePhoto="removePhotoFromList"
                    />
                </div>
                
                <div v-if="!isInMyBannedList && photosCount==0" class="mt-5 ">
                    <h2 class="d-flex justify-content-center" style="color: white;">No posts yet</h2>
                </div>
            </div>
        </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    
</template>

<style>
.profile-file-upload{
    display: none;
}
.my-btn-add-photo{
    background-color: green;
    border-color: grey;
}
.my-btn-add-photo:hover{
    color: white;
    background-color: green;
    border-color: grey;
}
.btn-foll{
    background-color: transparent;
    border: none;
    padding: 5px;
}
</style>
