import { Swiper, SwiperSlide } from "swiper/react";
import { Navigation, Pagination, A11y, Autoplay } from "swiper";
import './collectionswiper.styles.css';

const CollectionSwiper = () => {
  return (
    <>
      <Swiper
        modules={[Navigation, Pagination, A11y, Autoplay]}
        autoplay={{
          delay: 3000,
        }}
        navigation
        pagination={{ clickable: true }}
        breakpoints={{
          // when window width is >= 640px
          320: {
            slidesPerView: 2.8,
            spaceBetween: 50,
          },
          640: {
            slidesPerView: 4,
            spaceBetween: 50,
          },
          // when window width is >= 768px
          768: {
            slidesPerView: 5,
            spaceBetween: 50,
          },
          1200: {
            slidesPerView: 8,
            spaceBetween: 50,
          },
        }}
        onSlideChange={() => console.log("slide change")}
      >
        {[...Array(20)].map((val, key) => (
          <SwiperSlide className="swiper-slide">
            <div className="slide-div">
              <h1>{key}</h1>
            </div>
          </SwiperSlide>
        ))}
      </Swiper>
    </>
  );
};

export default CollectionSwiper;
