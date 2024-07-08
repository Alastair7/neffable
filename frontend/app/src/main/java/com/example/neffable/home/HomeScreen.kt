package com.example.neffable.home

import com.example.neffable.R
import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.material3.SnackbarHost
import androidx.compose.material3.SnackbarHostState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController

@Composable
fun HomeScreen(viewModel: HomeViewModel, navController: NavController) {
    val snackbarMessage by remember { viewModel.snackbarMessage }
    val snackbarHostState = remember { SnackbarHostState() }

    LaunchedEffect(snackbarMessage) {
        if (snackbarMessage.isNotEmpty()) {
            snackbarHostState.showSnackbar(snackbarMessage)
            viewModel.snackbarMessage.value = ""
        }
    }

    Box(modifier = Modifier.fillMaxSize()) {
        Column(
            modifier = Modifier.fillMaxSize(),
            verticalArrangement = Arrangement.SpaceBetween,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            // Icono superior centrado
            Image(
                painter = painterResource(id = R.drawable.he),
                contentDescription = "Top Icon",
                modifier = Modifier
                    .size(120.dp)
                    .padding(top = 24.dp)
            )
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.Center
            ) {
                Image(
                    painter = painterResource(id = R.drawable.love),
                    contentDescription = "Center Left Icon",
                    modifier = Modifier
                        .size(90.dp)
                        .padding(top = 24.dp)
                        .clickable { viewModel.sendEmotion("love") }
                )
                Image(
                    painter = painterResource(id = R.drawable.happy),
                    contentDescription = "Center Middle Icon",
                    modifier = Modifier
                        .size(120.dp)
                        .padding(start = 36.dp)
                        .clickable { viewModel.sendEmotion("happy") }
                )
                Image(
                    painter = painterResource(id = R.drawable.sad),
                    contentDescription = "Center Right Icon",
                    modifier = Modifier
                        .size(110.dp)
                        .padding(start = 48.dp)
                        .clickable { viewModel.sendEmotion("sad") }
                )
            }

            // Icono inferior centrado
            Image(
                painter = painterResource(id = R.drawable.she),
                contentDescription = "Bottom Icon",
                modifier = Modifier
                    .padding(bottom = 24.dp)
                    .size(90.dp)
            )
        }

        // Botón de regreso en la parte inferior izquierda
        Image(
            painter = painterResource(id = R.drawable.back_icon),
            contentDescription = "Back Icon",
            modifier = Modifier
                .size(50.dp)
                .padding(16.dp)
                .align(Alignment.BottomStart)
                .clickable { navController.navigate("mainPage") }
        )

        // SnackbarHost
        SnackbarHost(
            hostState = snackbarHostState,
            modifier = Modifier.align(Alignment.BottomCenter)
        )
    }
}
