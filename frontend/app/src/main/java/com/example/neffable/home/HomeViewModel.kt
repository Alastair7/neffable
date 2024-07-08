package com.example.neffable.home

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.coroutines.launch
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json

class HomeViewModel : ViewModel() {
    val snackbarMessage = mutableStateOf("")
    private val client = HttpClient {
        install(ContentNegotiation) {
            json(Json { ignoreUnknownKeys = true })
        }
    }

    fun sendEmotion(emotion: String) {
        viewModelScope.launch {
            try {
                val response: HttpResponse = client.post("http://localhost:5001/api/emotions") {
                    contentType(ContentType.Application.Json)
                    setBody(EmotionRequest(emotion))
                }
                if (response.status == HttpStatusCode.OK) {
                    val responseBody = response.body<BaseEmotionResponse>()
                    snackbarMessage.value = responseBody.message
                } else {
                    snackbarMessage.value = "Error: ${response.status}"
                }
            } catch (e: Exception) {
                snackbarMessage.value = "Error: ${e.message}"
            }
        }
    }

    override fun onCleared() {
        super.onCleared()
        client.close()
    }
}

@Serializable
data class EmotionRequest(val emotion: String)

@Serializable
data class BaseEmotionResponse(val statusCode: Int, val message: String, val emotion: String)
