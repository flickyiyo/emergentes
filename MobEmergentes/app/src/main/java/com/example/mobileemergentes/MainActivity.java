package com.example.mobileemergentes;

import androidx.appcompat.app.AppCompatActivity;

import android.graphics.Bitmap;
import android.graphics.BitmapFactory;
import android.os.Bundle;
import android.widget.ImageView;

import com.google.protobuf.ByteString;
import com.proto.imgstream.ImageRequest;
import com.proto.imgstream.ImageStream;
import com.proto.imgstream.ImgStreamServiceGrpc;

import java.util.Iterator;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        ImgStreamServiceGrpc service;
        ManagedChannel channel = ManagedChannelBuilder.forAddress("", 50051).build();
        ImgStreamServiceGrpc.ImgStreamServiceBlockingStub blockingStub = ImgStreamServiceGrpc.newBlockingStub(channel);
        ImgStreamServiceGrpc.ImgStreamServiceStub stub =
                ImgStreamServiceGrpc.newStub(channel);
        ImageRequest request = ImageRequest.newBuilder().setUsername("Edgar").build();
        Iterator<ImageStream> stream = blockingStub.askToRasppi(request);
        ImageView imv = findViewById(R.id.imv);
        while (stream.hasNext()) {
            ImageStream imageStream = stream.next();
            ByteString bytes = imageStream.getImage();
            Bitmap bmp = BitmapFactory.decodeFile(bytes.toString());
//            bmp.compress(Bitmap.CompressFormat.PNG, 100)
            imv.setImageBitmap(Bitmap.createBitmap(bmp));

        }


    }
}
