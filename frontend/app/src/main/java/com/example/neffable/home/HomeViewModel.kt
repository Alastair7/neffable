package com.example.neffable.home

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import io.ktor.client.*
import io.ktor.client.call.body
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.util.InternalAPI
import kotlinx.coroutines.launch

class HomeViewModel : ViewModel() {
    val snackbarMessage = mutableStateOf("")
    private val client = HttpClient()

    @OptIn(InternalAPI::class)
    fun sendEmotion(emotion: String) {
        viewModelScope.launch {
            try {
                val response: HttpResponse = client.post("http:localhost:5001/api/emotions") {
                    contentType(ContentType.Application.Json)
                    body = EmotionRequest(emotion)
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

data class EmotionRequest(val emotion: String)
data class BaseEmotionResponse(val statusCode: Int, val message: String, val emotion: String)
