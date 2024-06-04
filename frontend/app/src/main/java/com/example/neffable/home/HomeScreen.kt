package com.example.neffable.home
import com.example.neffable.R
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.viewModels
import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp

@Composable
fun HomeScreen(viewModel: MainViewModel) {
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

        // Iconos centrales en fila
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
            )
            Image(
                painter = painterResource(id = R.drawable.happy),
                contentDescription = "Center Middle Icon",
                modifier = Modifier
                    .size(120.dp)
                    .padding(start = 36.dp)
            )
            Image(
                painter = painterResource(id = R.drawable.sad),
                contentDescription = "Center Right Icon",
                modifier = Modifier
                    .size(110.dp)
                    .padding(start = 48.dp)
            )
        }

        // Icono inferior centrado
        Image(
            painter = painterResource(id = R.drawable.she),
            contentDescription = "Bottom Icon",
            modifier = Modifier.padding(bottom = 24.dp).size(90.dp)
        )
    }
}

