package com.example.neffable.home

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.neffable.data.RetrofitServiceFactory
import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json


class HomeViewModel : ViewModel() {
    val snackbarMessage = mutableStateOf("")
    val userConnectedResponseReceived = mutableStateOf(false)
    val isPulsing = mutableStateOf(false)

    private fun fetchPing() {
        CoroutineScope(Dispatchers.IO).launch {
            val service = RetrofitServiceFactory.makeRetrofitService()
            isPulsing.value = true  // Start pulsing animation

            try {
                val response = service.ping()
                if (response.isSuccessful) {
                    val pingResponse = response.body()
                    snackbarMessage.value = "Ping Response: ${pingResponse?.message}"
                    userConnectedResponseReceived.value = true
                } else {
                    snackbarMessage.value = "Error: ${response.code()} - ${response.message()}"
                }
            } catch (e: Exception) {
                snackbarMessage.value = "Exception: ${e.message}"
            } finally {
                isPulsing.value = false  // Stop pulsing animation
            }
        }
    }

    fun sendHappyEmotion(emotion: String) {
        snackbarMessage.value = "Sending happiness..."
        fetchPing()
    }

    fun sendSadEmotion(emotion: String) {
        snackbarMessage.value = "Sending sadness..."
        fetchPing()
    }

    fun sendLoveEmotion(emotion: String) {
        snackbarMessage.value = "Sending love..."
        fetchPing()
    }

    override fun onCleared() {
        super.onCleared()
    }
}



/*class HomeViewModel : ViewModel() {
    val snackbarMessage = mutableStateOf("")
    val userConnectedResponseReceived = mutableStateOf(false)
    val isPulsing = mutableStateOf(false)

    private fun fetchPing() {
        CoroutineScope(Dispatchers.IO).launch {
            val service = RetrofitServiceFactory.makeRetrofitService()
            isPulsing.value = true  // Start pulsing animation

            try {
                val response = service.ping()
                if (response.isSuccessful) {
                    val pingResponse = response.body()
                    snackbarMessage.value = "Ping Response: $pingResponse"
                    userConnectedResponseReceived.value = true
                } else {
                    snackbarMessage.value = "Error: ${response.code()} - ${response.message()}"
                }
            } catch (e: Exception) {
                snackbarMessage.value = "Exception: ${e.message}"
            } finally {
                isPulsing.value = false  // Stop pulsing animation
            }
        }
    }

    fun sendHappyEmotion(emotion: String) {
        snackbarMessage.value = "Sending happiness..."
        fetchPing()
    }

    fun sendSadEmotion(emotion: String) {
        snackbarMessage.value = "Sending sadness..."
        fetchPing()
    }

    fun sendLoveEmotion(emotion: String) {
        snackbarMessage.value = "Sending love..."
        fetchPing()
    }

    override fun onCleared() {
        super.onCleared()
    }
}*/

//class HomeViewModel : ViewModel() {
//    val snackbarMessage = mutableStateOf("")
//    private val client = HttpClient {
//        install(ContentNegotiation) {
//            json(Json { ignoreUnknownKeys = true })
//        }
//    }
//
//    val userConnectedResponseReceived = mutableStateOf(false)
//    val publicKey = "db674edc902017c71c8e0fdf6f22abd6"
//    val privateKey = "474e805f3bc50f23abc768bb68da6f7bbd7f8fe2"
//    var id = " "
//
//
//    fun sendHappyEmotion(emotion:String){
//        CoroutineScope(Dispatchers.IO).launch {
//            //snackbarMessage.value = "Sending happyness..."
//            id = "1016823"
//            //fetchHero(id, publicKey, privateKey)
//            //snackbarMessage.value = "Success!"
//            userConnectedResponseReceived.value = true
//        }
//    }
//    fun sendSadEmotion(emotion:String){
//        CoroutineScope(Dispatchers.IO).launch {
//            //snackbarMessage.value = "Sending sadness..."
//            id = "1017851"
//            //fetchHero(id, publicKey, privateKey)
//            //snackbarMessage.value = "Success!"
//            userConnectedResponseReceived.value = true
//        }
//    }
//    fun sendLoveEmotion(emotion: String) {
//        CoroutineScope(Dispatchers.IO).launch {
//            //snackbarMessage.value = "Sending love..."
//            id = "1017100"
//            //fetchHero(id, publicKey, privateKey)
//            //snackbarMessage.value = "Success!"
//            userConnectedResponseReceived.value = true
//        }
//        /*viewModelScope.launch {
//            try {
//                val service = RetrofitServiceFactory.makeRetrofitService()
//                val hero = service.getHero("db674edc902017c71c8e0fdf6f22abd6")
//                println(hero)
//                *//*val response: HttpResponse = client.post("http://localhost:5001/api/emotions") {
//                    contentType(ContentType.Application.Json)
//                    setBody(EmotionRequest(emotion))
////                }
////                if (response.status == HttpStatusCode.OK) {
//                    val responseBody = response.body<BaseEmotionResponse>()
//                    snackbarMessage.value = responseBody.message
//                } else {
//                    snackbarMessage.value = "Error: ${response.status}"
//                }*//*
//            } catch (e: Exception) {
//                snackbarMessage.value = "Error: ${e.message}"
//            }
//        }*/
//    }
//
//
//
//    override fun onCleared() {
//        super.onCleared()
//        client.close()
//    }
//}

@Serializable
data class EmotionRequest(val emotion: String)

@Serializable
data class BaseEmotionResponse(val statusCode: Int, val message: String, val emotion: String)
