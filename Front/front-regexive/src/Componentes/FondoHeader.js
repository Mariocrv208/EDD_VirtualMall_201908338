const ImageHeader = props => (
    <View style={{ backgroundColor: '#eee' }}>
      <Image
        style={StyleSheet.absoluteFill}
        source={{ uri: 'https://www.google.com/search?q=san+carlos+de+guatemala&tbm=isch&ved=2ahUKEwimxtu77crvAhVBWlkKHVk_Ba0Q2-cCegQIABAA&oq=san+carlos+de+guatemala&gs_lcp=CgNpbWcQAzICCAAyBggAEAUQHjIGCAAQBRAeMgYIABAFEB4yBggAEAUQHjIECAAQHjIECAAQHjIECAAQHjIGCAAQBRAeMgYIABAFEB46BAgjECc6BAgAEENQsd4DWJXyA2Cq8wNoAHAAeACAAasBiAG_EpIBBTEwLjEzmAEAoAEBqgELZ3dzLXdpei1pbWfAAQE&sclient=img&ei=4jJcYKbNN8G05QLZ_pToCg&bih=760&biw=1536#imgrc=mp-3_fktTsabiM' }}
      />
      <Header {...props} style={{ backgroundColor: 'transparent' }}/>
    </View>
  );